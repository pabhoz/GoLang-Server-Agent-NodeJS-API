import Model from '../libs/Model';

export default class Agent extends Model {

    static dbc: any;

    UID: string | undefined;

    constructor(agentUID: string) {
        super();
        this.UID = agentUID;
    }

    insert() {
        try {
            const query = `INSERT INTO servers.Agents (uid)
            VALUES ('${this.UID}');`;
            return super.execQuery(query);
        } catch (err) {
            return err;
        }
    }

    findOne() {
        try {
            const query = `SELECT * FROM servers.Agents WHERE uid = '${this.UID}';`;
            return super.execQuery(query);
        } catch (err) {
            return err;
        }
    }

}
