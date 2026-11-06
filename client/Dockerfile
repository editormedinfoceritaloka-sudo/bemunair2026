# Stage 1: full dependencies (build tooling included), used only to build
FROM node:22-alpine AS deps
WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN corepack enable && pnpm install --frozen-lockfile

# Stage 2: build
FROM node:22-alpine AS build
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build

# Stage 3: production-only dependencies, so devDependencies (vite, vitest,
# typescript, eslint, ...) never ship in the runtime image.
FROM node:22-alpine AS prod-deps
WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN corepack enable && pnpm install --prod --frozen-lockfile

# Stage 4: runner
FROM node:22-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
ENV PORT=3000
RUN chown node:node /app
COPY --chown=node:node --from=build /app/build ./build
COPY --chown=node:node --from=build /app/package.json ./package.json
COPY --chown=node:node --from=prod-deps /app/node_modules ./node_modules
USER node
EXPOSE 3000
HEALTHCHECK --interval=15s --timeout=5s --start-period=20s --retries=10 \
  CMD node -e "fetch('http://127.0.0.1:3000/admin/login').then(r=>process.exit(r.ok?0:1)).catch(()=>process.exit(1))"
CMD ["node", "build/index.js"]
