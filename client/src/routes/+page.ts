// +page.server.ts fetches live cabinet/article data from the backend API on
// every request, so this route cannot be statically prerendered: doing so
// makes `npm run build` (and therefore `docker build`) call that API at
// build time, which fails whenever the API isn't reachable from the build
// environment (CI, a standalone `docker build`, ...) and bakes stale data
// into the output the rest of the time. Render it per-request instead.
export const prerender = false;
