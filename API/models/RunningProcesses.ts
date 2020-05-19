import Model from '../libs/Model';

interface ProcessInfo {
    pid: number
    name: string
}

export default class RunningProcesses extends Model {

    static dbc: any;

    agentUID: string | undefined;
    total: string;
    running: string;
    processesList: Array<ProcessInfo>;

    constructor(total: string,
        running: string,
        processesList: Array<ProcessInfo>) {
        super();
        this.total = total;
        this.running = running;
        this.processesList = processesList;
    }

    setAgentUID(agentUID: string) {
        this.agentUID = agentUID;
    }

    insert() {
        try {
            if (this.agentUID === undefined) { throw "No Agent UID assigned"; }
            const query = `INSERT INTO servers.RunningProcessesLog (agentId, total, running, processesList)
            VALUES ('${this.agentUID}', '${this.total}', '${this.running}', '${JSON.stringify(this.processesList)}');`;
            return super.insert(query);
        } catch (err) {
            return err;
        }
    }

}
