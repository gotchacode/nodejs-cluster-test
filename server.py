from aiohttp import web


async def handle(request):
    return web.Response(text='Hello World')


app = web.Application()
app.router.add_get('/', handle)
web.run_app(app, host='127.0.0.1', port=8000)

