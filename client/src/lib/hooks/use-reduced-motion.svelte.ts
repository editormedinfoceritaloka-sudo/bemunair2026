import { onMount } from 'svelte';

export function useReducedMotion() {
  let reduced = $state(false);

  onMount(() => {
    const query = window.matchMedia('(prefers-reduced-motion: reduce)');
    const update = () => { reduced = query.matches; };
    update();
    query.addEventListener('change', update);
    return () => query.removeEventListener('change', update);
  });

  return { get value() { return reduced; } };
}
