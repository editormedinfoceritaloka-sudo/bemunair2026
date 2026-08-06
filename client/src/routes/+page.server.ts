import { requireApiData } from '$lib/server/api';
import type { Cabinet } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
  const cabinet = await requireApiData<Cabinet>(
    fetch,
    undefined,
    '/cabinet'
  );
  return {
    cabinet
  };
};