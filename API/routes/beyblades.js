module.exports = function(app){

    app.get('/bayblades', (req, res) => {
        res.send(`[GET] ${req.url}`);
    });
    
    app.post('/bayblades', (req, res) => {
        res.status(200).json({ bayblades: [] });
    });
    
    app.put('/bayblades', (req, res) => {
        res.send(`[PUT] ${req.url}`);
    });
    
    app.delete('/bayblades', (req, res) => {
        res.send(`[DELETE] ${req.url}`);
    });

}


