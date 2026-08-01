import { apiRequest, tokenFrom } from '$lib/server/api';
import type { ContentSubmission } from '$lib/types';
import type { PageServerLoad } from './$types';

type SubmissionType = 'CONTENT' | 'ARTICLE';

export const load: PageServerLoad = async ({ fetch, cookies, url }) => {
  const type: SubmissionType = url.searchParams.get('type') === 'ARTICLE' ? 'ARTICLE' : 'CONTENT';
  const { data: submissions } = await apiRequest<ContentSubmission[]>(
    fetch,
    tokenFrom(cookies),
    '/content-submissions'
  );

  return {
    type,
    submissions: submissions.filter((submission) => submission.service_type === type)
  };
};
