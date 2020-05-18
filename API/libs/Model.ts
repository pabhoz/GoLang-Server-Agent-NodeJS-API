const mysql = require('mysql');

export default class Model {

    static dbc: any = mysql.createConnection({
        host: 'localhost',
        user: 'root',
        password: ''
    }); 
    
    constructor() { }
    
    insert(query: string) {
        let result = new Promise((resolve, reject) => {
            Model.dbc.query(query, function (err: any, result: any) {
                if (err) reject(err);
                resolve(result)
              });
        });

        return result;
    }
}
