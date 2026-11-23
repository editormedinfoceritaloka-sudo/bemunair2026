import type { SubmitFunction } from '@sveltejs/kit';
import { toast } from 'svelte-sonner';

export function useFormSubmit(message: string, onSuccess?: () => void): SubmitFunction {
  return async () => async ({ result, update }) => {
    await update();
    if (result.type !== 'success') return;
    onSuccess?.();
    toast.success(message);
  };
}
