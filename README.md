# Pi Slide API

This is a really early API (and eventually a UI) for managing various aspects of [pislide-os](https://github.com/JarvyJ/pislide-os). It can currently do the following:
- Change slideshow settings
- Start/Stop the slideshow service (not in the UI yet)
- Upload and manage photos (not in the UI yet)

The UI is definitely not ready for full-time use, but the API can be tested via the `/api-ref` endpoint.

## Other planned features
- Shutdown/Restart the OS
- Resize large images on upload (if needed)
- Convert small looping videos into actual gifs
- Finish and zhuzh up the frontend

## Frontend
The frontend is written with [SvelteKit](https://kit.svelte.dev/) and uses [Bulma](https://bulma.io/) for CSS. It's built statically and gets bundled in `frontend/build`. It's using the current LTS version of node (as of this writing: node v20.15 - npm v10.7.0) and I'm trying out Svelte 5.

## Building and Running
Assuming you have golang setup, after pulling down this repository:
```bash
go run cmd/pislide-api/main.go
```
(or run with [gow](https://github.com/mitranim/gow) for hot reloading)

The API will be available at [http://localhost:8888/api-ref](http://localhost:8888/api-ref). It will run the frontend code found in `frontend/build`.

If you are working on frontend development, you'll want to start the frontend process:
```bash
cd frontend
npm run dev -- --open
```
and then go to [http://localhost:5173/](http://localhost:5173/) to see the live changes.
