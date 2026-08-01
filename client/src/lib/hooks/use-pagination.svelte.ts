export function pageFromSearch(search: URLSearchParams) {
  const value = Number(search.get('page') || '1');
  return Number.isFinite(value) && value > 0 ? Math.floor(value) : 1;
}

export function paginationPages(page: number, totalPages: number) {
  const first = Math.max(1, page - 2);
  const last = Math.min(totalPages, first + 4);
  return Array.from({ length: Math.max(0, last - first + 1) }, (_, index) => first + index);
}
