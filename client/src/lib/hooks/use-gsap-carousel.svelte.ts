import { gsap } from 'gsap';

export function useGsapCarousel() {
  let context: gsap.Context | null = null;

  function animate(root: HTMLElement, activeIndex: number, total: number, reduced: boolean) {
    const cards = Array.from(root.querySelectorAll<HTMLElement>('[data-carousel-card]'));
    context?.revert();
    context = gsap.context(() => {
      cards.forEach((card, index) => {
        const offset = (index - activeIndex + total) % total;
        const normalized = offset > total / 2 ? offset - total : offset;
        const distance = Math.min(Math.abs(normalized), 2);
        const visible = distance <= 2;
        gsap.to(card, { x: normalized * 190, scale: distance === 0 ? 1 : 0.78, opacity: visible ? (distance === 0 ? 1 : 0.5) : 0, zIndex: distance === 0 ? 3 : 2, duration: reduced ? 0 : 0.4, ease: 'power3.out', overwrite: true });
      });
    }, root);
  }

  function destroy() {
    context?.revert();
    context = null;
  }

  return { animate, destroy };
}
