# Coforge Interseguro Challenge

Implementación de un servicio de procesamiento de matrices compuesto por dos microservicios:

- `ms-go`: expone la API principal, rota la matriz y calcula su factorización QR.
- `ms-node`: calcula estadísticas sobre matrices y es consumido por `ms-go`.

## Arquitectura

```text
Cliente
	|
	| POST /login
	| POST /matrix + JWT
	v
ms-go :3000
	|
	| POST /api/get-statistics
	v
ms-node :3001
```

Para cada solicitud a `ms-go`, el servicio:

1. Valida la matriz y comprueba que sea rectangular.
2. La rota 90 grados en sentido horario.
3. Calcula la factorización QR de la matriz rotada.
4. Solicita las estadísticas de la matriz rotada, `Q` y `R` a `ms-node` en paralelo.
5. Devuelve las matrices y sus estadísticas.

## Requisitos

- Docker y Docker Compose para ejecutar todo el sistema.
- Go `1.26.2` o compatible para ejecutar `ms-go` localmente.
- Node.js `24` y pnpm `10` para ejecutar `ms-node` localmente.

## Ejecución con Docker Compose

Desde la raíz del repositorio:

```bash
docker compose up --build
```

Los servicios quedan disponibles en:

- API principal: `http://localhost:3000`
- API de estadísticas: `http://localhost:3001`

Para detenerlos:

```bash
docker compose down
```

## Ejecución local

### Servicio Go

```bash
cd ms-go
export PORT=3000
export STATISTICS_API_URL=http://localhost:3001/api
export JWT_SECRET=local-development-secret
go run ./cmd/api
```

### Servicio Node

En otra terminal:

```bash
cd ms-node
pnpm install
PORT=3001 pnpm dev
```

`ms-node` usa el puerto `3001` por defecto. `ms-go` requiere definir `PORT` y `STATISTICS_API_URL`.

## API

### `POST /login`

Autentica un usuario y devuelve un JWT válido durante una hora. Este endpoint es público y está servido por `ms-go`.

Solicitud:

```bash
curl -X POST http://localhost:3000/login \
	-H 'Content-Type: application/json' \
	-d '{"username":"admin","password":"password"}'
```

Respuesta `200 OK`:

```json
{
	"token": "eyJ..."
}
```

Las credenciales de prueba son `admin` y `password`. Las credenciales inválidas se responden con estado `401 Unauthorized`.

### `POST /matrix`

Endpoint principal, servido por `ms-go`. Requiere un JWT en la cabecera `Authorization` con el formato `Bearer <token>`.

Solicitud:

```bash
curl -X POST http://localhost:3000/matrix \
	-H 'Content-Type: application/json' \
	-H 'Authorization: Bearer <token-obtenido-en-login>' \
	-d '{"matrix":[[1,2,3],[4,5,6]]}'
```

Las solicitudes sin un JWT válido se responden con estado `401 Unauthorized`.

Respuesta `200 OK`:

```json
{
	"rotatedMatrix": [[4, 1], [5, 2], [6, 3]],
	"rotatedStatistics": {
		"max": 6,
		"min": 1,
		"mean": 3.5,
		"sum": 21,
		"isDiagonal": false
	},
	"q": [[0.455842, -0.790912], [0.569803, -0.093048], [0.683763, 0.604815]],
	"qStatistics": {
		"max": 0.683763,
		"min": -0.790912,
		"mean": 0.238377,
		"sum": 1.430263,
		"isDiagonal": false
	},
	"r": [[8.774964, 3.646738], [0, 0.837436]],
	"rStatistics": {
		"max": 8.774964,
		"min": 0,
		"mean": 3.117285,
		"sum": 12.469138,
		"isDiagonal": false
	}
}
```

Los valores de `q` y `r` pueden variar ligeramente por redondeo de punto flotante.

La matriz de entrada debe ser no vacía y rectangular. Además, la matriz rotada debe cumplir las dimensiones necesarias para la factorización QR y no tener columnas linealmente dependientes. Los errores de validación se responden con estado `400`.

### `POST /api/get-statistics`

Endpoint interno, servido por `ms-node`.

Solicitud:

```bash
curl -X POST http://localhost:3001/api/get-statistics \
	-H 'Content-Type: application/json' \
	-d '{"matrix":[[1,2],[3,4]]}'
```

Respuesta `200 OK`:

```json
{
	"max": 4,
	"min": 1,
	"mean": 2.5,
	"sum": 10,
	"isDiagonal": false
}
```

`matrix` debe ser un arreglo no vacío de arreglos numéricos.

### `GET /health-check`

Comprobación de disponibilidad de `ms-node`:

```bash
curl http://localhost:3001/health-check
```

Respuesta: `[MS-NODE] Health check passed!`

## Pruebas y calidad

Ejecutar las pruebas de Go:

```bash
cd ms-go
go test ./...
```

Ejecutar las pruebas de Node.js:

```bash
cd ms-node
pnpm install
pnpm test
```

El chequeo y formateo de Node.js se ejecuta con:

```bash
cd ms-node
pnpm check
```

## Variables de entorno

| Servicio | Variable | Valor | Descripción |
| --- | --- | --- | --- |
| `ms-go` | `PORT` | `3000` en Compose | Puerto HTTP de la API principal. |
| `ms-go` | `STATISTICS_API_URL` | `http://ms-node:3001/api` en Compose | URL base usada para consultar estadísticas. |
| `ms-go` | `JWT_SECRET` | `local-development-secret` localmente | Secreto usado para firmar y validar los JWT. Debe configurarse con un valor seguro en otros entornos. |
| `ms-node` | `PORT` | `3001` | Puerto HTTP de la API de estadísticas. |