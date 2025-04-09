
htmx.on('htmx:configRequest', (event) => event.detail.path = `/partial${event.detail.path}`);
