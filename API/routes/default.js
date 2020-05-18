module.exports = function (app) {
    app.get('*', (req, res) => {
        res.send(`[GET] ${req.url}`);
    });

    app.post('*', (req, res) => {
        res.status(200).json({ hello: "world" });
    });

    app.put('*', (req, res) => {
        res.send(`[PUT] ${req.url}`);
    });

    app.delete('*', (req, res) => {
        res.send(`[DELETE] ${req.url}`);
    });
}
