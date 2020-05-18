module.exports = function(app, db){

    app.get('/gods', (req, res) => {
        res.status(200).json(db.gods);
    });
    app.get('/gods/:name', (req, res) => {
        const name = req.params.name;
        const god = db.gods.filter(el => {
            return el.name == name;
        })
        res.status(200).json(god);
    });

    app.get('/gods/origin/:origin', (req, res) => {
        const origin = req.params.origin;
        const god = db.gods.filter(el => {
            return el.origin == origin;
        })
        res.status(200).json(god);
    });
    
    app.post('/gods', (req, res) => {
        const god = {
            name: req.body.name,
            origin: req.body.origin
        };
        db.gods.push(god);
        res.status(201).json({msg: "Created", result: god});
    });
    
    app.put('/gods', (req, res) => {
        res.send(`[PUT] ${req.url}`);
    });
    
    app.delete('/gods', (req, res) => {
        res.send(`[DELETE] ${req.url}`);
    });

}


