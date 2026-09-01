dev-back:
	cd backend && air

dev-front:
	cd frontend && npm run dev

dev:
	make -j2 dev-back dev-front
