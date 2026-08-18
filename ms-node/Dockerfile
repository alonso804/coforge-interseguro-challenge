FROM node:24-alpine AS base

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN corepack enable && \
    corepack prepare pnpm@10 --activate

WORKDIR /app


FROM base AS dependencies

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

RUN pnpm install --frozen-lockfile


FROM base AS build

COPY --from=dependencies /app/node_modules ./node_modules
COPY . .

RUN pnpm build


FROM node:24-alpine AS production

ENV NODE_ENV=production

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN corepack enable && \
    corepack prepare pnpm@10 --activate && \
    pnpm install --frozen-lockfile --prod

COPY --from=build /app/dist ./dist

EXPOSE 3001

CMD ["node", "dist/index.mjs"]
