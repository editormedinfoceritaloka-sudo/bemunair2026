FROM node:22-bookworm-slim

WORKDIR /app

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV SHARP_IGNORE_GLOBAL_LIBVIPS=1

RUN corepack enable

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./
COPY .npmrc* ./

RUN pnpm install --frozen-lockfile

COPY . .

ENV NODE_ENV=production

RUN mkdir -p /app/sessions && chown -R node:node /app
USER node

EXPOSE 3001

CMD ["pnpm", "start"]
