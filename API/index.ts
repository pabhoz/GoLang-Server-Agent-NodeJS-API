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

app.use(function (req, res, next) {
    const agentKey = req.header("Agent-Key");
    if (!req.header("Agent-Key")) {
      return res.status(403).json({ error: 'No credentials sent!' });
    } else {
        if (agentKey !== "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918") {
            return res.status(401).json({error: "Not an authorized agent"})
        }
    }
    next();
});
  
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
