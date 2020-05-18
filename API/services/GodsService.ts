import God from "../models/God";

const { db } = require('../conf');

module GodsService {
    export function get(id: number): God {
        const god = db.gods.filter((el: God) => {
            return el.id == id;
        })
        return god;
    }

    export function getByName(name: string): God {
        const god = db.gods.filter((el: God) => {
            return el.name == name;
        })
        return god;
    }

    export function getAll(): Promise<any> {

        let result = new Promise((resolve, reject) => {
            God.dbc.query('SELECT * from rest.god', (err: any, rows: any, fields: any) => {
                if (err) { reject(err); }
                resolve(rows);
            });
        });

        return result;
    }

    export function create(name: string, origin: string): God {
        const god = new God(name, origin);
        db.gods.push(god);
        return god;
    }
}

export default GodsService;
