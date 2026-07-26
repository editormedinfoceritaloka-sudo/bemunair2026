import { clearSession } from '$lib/server/api';
import { redirect } from '@sveltejs/kit';
export const POST = ({ cookies }) => { clearSession(cookies); throw redirect(303, '/admin/login'); };
