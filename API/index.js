const express = require("express");
const bodyParser = require("body-parser");
const app = express();
const port = 8888;

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.post('/processor', (req, res) => {
    res.status(200).json({ msg: "you rock" });
});

// Iniciamos la escucha de peticiones por el puerto 3000
app.listen(port, () => {
    console.log(`Node JS Server started at port ${port}`);
});
