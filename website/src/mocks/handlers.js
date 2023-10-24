// src/mocks/handlers.js
import { rest } from 'msw'

let clusters = [
    {
        code: "АСТАНА",
        cluster: "КЗ",
        source: "WB"
    }, {
        code: "БЕЛАЯ ДАЧА",
        cluster: "Мск",
        source: "WB"
    }, {
        code: "ВЁШКИ",
        cluster: "Мск",
        source: "WB"
    }, {
        code: "ДОМОДЕДОВО",
        cluster: "Мск",
        source: "WB"
    }
];

export const handlers = [

    rest.post('/rest/clusters', async (req, res, ctx) => {
         const record = await req.json()
         clusters = [{ ...record}, ...clusters]
        return res(
            // Respond with a 200 status code
            ctx.status(200),
            ctx.json(record)
        );
    }),

    rest.put('/rest/clusters', async (req, res, ctx) => {
        const record = await req.json()
        clusters = clusters.map(item => {
            if (item.code === record.code) {
                return record
            }
            return item
        })
        return res(
            // Respond with a 200 status code
            ctx.status(200),
            ctx.json(record)
        );
    }),

    rest.get('/rest/clusters', (req, res, ctx) => {
        const page = req.params["page"] ?? 1
        return res(
            ctx.status(200),
            ctx.json({
                page: page,
                size: 50,
                count: clusters.length,
                data: clusters.slice((page - 1) * 50, (page) * 50)
            }),
        )
    }),

    // rest.get('/rest/sales/report', (req, res, ctx) => {
    //     const year = req.params['year']
    //     const month = req.params['month']
    //     const content = "Hello!"
    //     const blob = new Blob([content], { type: 'text/csv' });
    //     return res(
    //         ctx.status(200),
    //         ctx.set({
    //             'Content-Disposition': 'attachment; filename="things_export.csv"',
    //             'Content-type': 'text/csv',
    //             'Content-Length': blob.size.toString(),
    //         }),
    //         ctx.body(blob),
    //         ctx.delay(),
    //     )
    //
    // })
]