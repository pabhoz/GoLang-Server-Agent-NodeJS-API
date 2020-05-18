import express from "express";
import bodyParser from "body-parser";
const app = express();

const port = 3000;

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

import * as ProcessorsController from "./controllers/ProcessorsController";

app.route('/processors')
.post(ProcessorsController.createLog)

app.listen(port, () => {
    console.log(`Node JS Server started at port ${port}`);
});
