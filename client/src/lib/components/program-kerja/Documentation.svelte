<script lang="ts">
  type SlotId =
    | 'a'
    | 'b'
    | 'c'
    | 'd'
    | 'e'
    | 'f';

  let {
    images,
    title = 'Dokumentasi'
  }: {
    images: string[];
    title?: string;
  } = $props();

  const slots: SlotId[] = [
    'a',
    'b',
    'c',
    'd',
    'e',
    'f'
  ];

  let failedSlots = $state<
    Record<SlotId, boolean>
  >({
    a: false,
    b: false,
    c: false,
    d: false,
    e: false,
    f: false
  });

  const displayedImages = $derived.by(() => {
    return images.slice(0, 6);
  });

  $effect(() => {
    displayedImages.join('|');

    failedSlots = {
      a: false,
      b: false,
      c: false,
      d: false,
      e: false,
      f: false
    };
  });

  function handleImageError(
    slot: SlotId
  ): void {
    failedSlots[slot] = true;
  }
</script>

<section class="w-full">
  <h2
    class="
      text-center
      text-4xl leading-none font-black
      tracking-[-0.055em]
      text-orange-600
      [-webkit-text-stroke:2px_white]
      [paint-order:stroke_fill]
      drop-shadow-[0_6px_2px_rgba(15,23,42,0.32)]
      sm:text-5xl
      lg:text-6xl
    "
  >
    {title}
  </h2>

  <div
    class="
      mt-10 w-full
      overflow-hidden
      rounded-[24px]
      bg-blue-950
      p-4
      shadow-[0_18px_35px_rgba(5,34,70,0.24)]
      sm:p-6
    "
  >
    <div class="documentation-grid">
      {#each slots as slot, index (slot)}
        {@const imageSrc = displayedImages[index]}

        <figure
          class={`documentation-slot slot-${slot}`}
        >
          {#if imageSrc && !failedSlots[slot]}
            <img
              src={imageSrc}
              alt={`Dokumentasi program kerja ${index + 1}`}
              loading="lazy"
              decoding="async"
              onerror={() => {
                handleImageError(slot);
              }}
              class="
                h-full w-full
                object-cover object-center
                transition-transform
                duration-500
                hover:scale-[1.03]
              "
            />
          {:else}
            <div
              class="documentation-placeholder"
              role="img"
              aria-label="Dokumentasi belum tersedia"
            ></div>
          {/if}
        </figure>
      {/each}
    </div>
  </div>
</section>

<style>
  .documentation-grid {
    display: grid;
    grid-template-columns: 1fr;
    grid-auto-rows: 190px;
    gap: 0.75rem;
  }

  .documentation-slot {
    position: relative;
    min-width: 0;
    min-height: 0;
    overflow: hidden;
    border-radius: 18px;
    background-color: white;
  }

  .documentation-placeholder {
    width: 100%;
    height: 100%;
    background-color: #f8fafc;
    background-image:
      linear-gradient(
        45deg,
        #e5e7eb 25%,
        transparent 25%
      ),
      linear-gradient(
        -45deg,
        #e5e7eb 25%,
        transparent 25%
      ),
      linear-gradient(
        45deg,
        transparent 75%,
        #e5e7eb 75%
      ),
      linear-gradient(
        -45deg,
        transparent 75%,
        #e5e7eb 75%
      );
    background-size: 32px 32px;
    background-position:
      0 0,
      0 16px,
      16px -16px,
      -16px 0;
  }

  @media (min-width: 640px) {
    .documentation-grid {
      grid-template-columns:
        repeat(2, minmax(0, 1fr));
      grid-auto-rows: 210px;
    }

    .slot-c {
      grid-row: span 2;
    }
  }

  @media (min-width: 1024px) {
    .documentation-grid {
      height: clamp(
        430px,
        48vw,
        560px
      );
      grid-template-columns:
        repeat(12, minmax(0, 1fr));
      grid-template-rows:
        repeat(7, minmax(0, 1fr));
      grid-auto-rows: auto;
      gap: 0.75rem;
    }

    .slot-a {
      grid-column: 1 / 6;
      grid-row: 1 / 4;
    }

    .slot-b {
      grid-column: 6 / 10;
      grid-row: 1 / 4;
    }

    .slot-c {
      grid-column: 10 / 13;
      grid-row: 1 / 5;
    }

    .slot-d {
      grid-column: 1 / 5;
      grid-row: 4 / 8;
    }

    .slot-e {
      grid-column: 5 / 10;
      grid-row: 4 / 8;
    }

    .slot-f {
      grid-column: 10 / 13;
      grid-row: 5 / 8;
    }
  }
</style>