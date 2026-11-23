type AudioContextConstructor = typeof AudioContext;

type SoundName =
  | 'book'
  | 'page'
  | 'flip'
  | 'swipe'
  | 'type'
  | 'shine';

const SOUND_URLS: Record<SoundName, string> = {
  book: '/audio/book-open.mp3',
  page: '/audio/page-turn.mp3',
  flip: '/audio/flip-card.mp3',
  swipe: '/audio/swipe.mp3',
  type: '/audio/keyboard-typing.mp3',
  shine: '/audio/shine.mp3'
};

type QueuedSound = {
  name: SoundName;
  volume: number;
  rate: number;
};

let context: AudioContext | null = null;
let unlocked = false;

const buffers = new Map<SoundName, AudioBuffer>();
const pending = new Map<SoundName, Promise<AudioBuffer | null>>();
const queuedSounds: QueuedSound[] = [];

let lastTypeSound = 0;

function getContext(): AudioContext | null {
  if (typeof window === 'undefined') return null;

  const AudioContextClass: AudioContextConstructor | undefined =
    window.AudioContext ??
    (window as Window & { webkitAudioContext?: AudioContextConstructor })
      .webkitAudioContext;

  if (!AudioContextClass) return null;

  context ??= new AudioContextClass();

  if (context.state === 'suspended') {
    void context.resume().catch(() => {});
  }

  return context;
}

function unlock(): void {
  if (unlocked || typeof window === 'undefined') return;
  unlocked = true;
  getContext();

  const queued = queuedSounds.splice(0);
  queued.forEach(({ name, volume, rate }) => {
    play(name, volume, rate);
  });
}

function loadBuffer(name: SoundName): Promise<AudioBuffer | null> {
  const cached = buffers.get(name);
  if (cached) return Promise.resolve(cached);

  const inFlight = pending.get(name);
  if (inFlight) return inFlight;

  const ctx = getContext();
  if (!ctx) return Promise.resolve(null);

  const task = fetch(SOUND_URLS[name])
    .then((response) => {
      if (!response.ok) throw new Error(String(response.status));
      return response.arrayBuffer();
    })
    .then((data) => ctx.decodeAudioData(data))
    .then((buffer) => {
      buffers.set(name, buffer);
      pending.delete(name);
      return buffer;
    })
    .catch(() => {
      pending.delete(name);
      return null;
    });

  pending.set(name, task);
  return task;
}

function play(name: SoundName, volume: number, rate = 1): void {
  const ctx = getContext();
  if (!ctx || document.hidden) return;

  if (!unlocked) {
    return;
  }

  void loadBuffer(name).then((buffer) => {
    if (!buffer || document.hidden) return;

    const source = ctx.createBufferSource();
    source.buffer = buffer;
    source.playbackRate.value = rate;

    const gain = ctx.createGain();
    gain.gain.value = volume;

    source.connect(gain).connect(ctx.destination);
    source.start();
  });
}

if (typeof window !== 'undefined') {
  window.addEventListener('pointerdown', unlock, {
    once: true,
    passive: true
  });
  window.addEventListener('wheel', unlock, {
    once: true,
    passive: true
  });
  window.addEventListener('touchstart', unlock, {
    once: true,
    passive: true
  });
  window.addEventListener('keydown', unlock, { once: true });

  (Object.keys(SOUND_URLS) as SoundName[]).forEach(loadBuffer);
}

export function playBookSound(): void {
  play('book', 1);
}

export function playPageSound(): void {
  play('page', 0.95);
}

export function playFlipSound(): void {
  play('flip', 0.55);
}

export function playSwipeSound(): void {
  play('swipe', 1);
}

export function playShineSound(): void {
  play('shine', 0.45);
}

export function playTypeSound(): void {
  const ctx = getContext();
  if (!ctx || ctx.state !== 'running') return;

  const now = performance.now();
  if (now - lastTypeSound < 90) return;
  lastTypeSound = now;

  void loadBuffer('type').then((buffer) => {
    if (!buffer) return;

    const source = ctx.createBufferSource();
    source.buffer = buffer;
    source.playbackRate.value =
      0.86 + Math.random() * 0.1;

    const gain = ctx.createGain();
    const duration = 0.07 + Math.random() * 0.045;
    const startAt = ctx.currentTime;
    const offset =
      Math.random() *
      Math.max(0, buffer.duration - duration - 0.05);

    gain.gain.setValueAtTime(0.0001, startAt);
    gain.gain.linearRampToValueAtTime(0.18, startAt + 0.006);
    gain.gain.setValueAtTime(
      0.18,
      startAt + Math.max(0.006, duration - 0.012)
    );
    gain.gain.linearRampToValueAtTime(
      0.0001,
      startAt + duration
    );

    source.connect(gain).connect(ctx.destination);
    source.start(startAt, offset, duration);
  });
}
