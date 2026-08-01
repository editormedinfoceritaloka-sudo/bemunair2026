import { invalidateAll } from '$app/navigation';
import type { SubmitFunction } from '@sveltejs/kit';
import { toast } from 'svelte-sonner';

export function useFormSubmit(message: string, onSuccess?: () => void): SubmitFunction {
  return async () => async ({ update }) => {
    await update();
    await invalidateAll();
    onSuccess?.();
    toast.success(message);
  };
}
