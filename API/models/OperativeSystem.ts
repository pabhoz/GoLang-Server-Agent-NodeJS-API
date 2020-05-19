import Model from '../libs/Model';

export default class OperativeSystem extends Model {

    static dbc: any;

    agentUID: string | undefined;
    runtime: string;
    name: string;
    platform: string;

    constructor(runtime: string,
        name: string,
        platform: string) {
        super();
        this.runtime = runtime;
        this.name = name;
        this.platform = platform;
    }

    setAgentUID(agentUID: string) {
        this.agentUID = agentUID;
    }

    insert() {
        try {
            if (this.agentUID === undefined) { throw "No Agent UID assigned"; }
            const query = `INSERT INTO servers.OSLogs (agentId, runtime, name, platform)
            VALUES ('${this.agentUID}', '${this.runtime}', '${this.name}', '${this.platform}');`;
            //console.log(query);
            return super.execQuery(query);
        } catch (err) {
            return err;
        }
    }

}
