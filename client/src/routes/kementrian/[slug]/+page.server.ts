import { error } from '@sveltejs/kit';
import { apiRequest } from '$lib/server/api';
import type {
  OrganizationUnit,
  WorkProgram
} from '$lib/types';
import type { PageServerLoad } from './$types';

function getErrorStatus(cause: unknown): number {
  if (
    typeof cause === 'object' &&
    cause !== null &&
    'status' in cause &&
    typeof cause.status === 'number'
  ) {
    return cause.status;
  }

  return 500;
}

export const load: PageServerLoad = async ({
  fetch,
  params
}) => {
  try {
    const [
      unitResponse,
      programsResponse
    ] = await Promise.all([
      apiRequest<OrganizationUnit>(
        fetch,
        undefined,
        `/cabinet/units/${params.slug}`
      ),
      apiRequest<WorkProgram[]>(
        fetch,
        undefined,
        `/cabinet/units/${params.slug}/programs?page=1`
      )
    ]);

    return {
      unit: unitResponse.data,
      programs: programsResponse.data
    };
  } catch (cause) {
    console.error(
      `Gagal mengambil kementerian "${params.slug}":`,
      cause
    );

    const status = getErrorStatus(cause);

    if (status === 404) {
      error(
        404,
        'Kementerian tidak ditemukan'
      );
    }

    if (status === 503) {
      error(
        503,
        'Server API tidak dapat dihubungi'
      );
    }

    error(
      status,
      'Gagal mengambil data kementerian'
    );
  }
};