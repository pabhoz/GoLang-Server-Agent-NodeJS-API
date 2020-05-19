import Model from '../libs/Model';

export default class Users extends Model {

    static dbc: any;

    agentUID: string | undefined;
    activeUsers: string;

    constructor(activeUsers: string) {
        super();
        this.activeUsers = activeUsers;
    }

    setAgentUID(agentUID: string) {
        this.agentUID = agentUID;
    }

    insert() {
        try {
            if (this.agentUID === undefined) { throw "No Agent UID assigned"; }
            const query = `INSERT INTO servers.UsersLogs (agentId, activeUsers)
            VALUES ('${this.agentUID}', '${JSON.stringify(this.activeUsers)}');`;
            return super.execQuery(query);
        } catch (err) {
            return err;
        }
    }

}
