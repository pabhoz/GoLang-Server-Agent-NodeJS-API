import express from "express";
import bodyParser from "body-parser";
const app = express();

const port = 3000;

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

import * as AgentsController from "./controllers/AgentsController";
import * as ProcessorsController from "./controllers/ProcessorsController";
import * as RunningProcessesController from "./controllers/RunningProcessesController";
import * as UsersController from "./controllers/UsersController";
import * as OperativeSystemsController from "./controllers/OperativeSystemsController";

app.route('/agents')
    .post(AgentsController.handshake)

app.route('/processors')
    .post(ProcessorsController.createLog)

app.route('/runningProcesses')
    .post(RunningProcessesController.createLog)

app.route('/users')
    .post(UsersController.createLog)

app.route('/os')
    .post(OperativeSystemsController.createLog)

app.listen(port, () => {
    console.log(`Node JS Server started at port ${port}`);
});
