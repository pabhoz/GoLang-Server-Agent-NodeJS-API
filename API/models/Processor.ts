import Model from '../libs/Model';

interface CPUUtilizationI {
    index: number,
    utilization: string
}

export default class Processor extends Model {

    static dbc: any;

    agentUID: string | undefined;
    cpuIndex: number;
    vendorId: string;
    family: string;
    numberOfCores: number;
    modelName: string;
    speed: string;
    currentCPUUtilization: Array<CPUUtilizationI>

    constructor(cpuIndex: number,
        vendorId: string,
        family: string,
        numberOfCores: number,
        modelName: string,
        speed: string,
        currentCPUUtilization: Array<CPUUtilizationI>) {
        super();
        this.cpuIndex = cpuIndex;
        this.vendorId = vendorId;
        this.family = family;
        this.numberOfCores = numberOfCores;
        this.modelName = modelName;
        this.speed = speed;
        this.currentCPUUtilization = currentCPUUtilization;
    }

    setAgentUID(agentUID: string) {
        this.agentUID = agentUID;
    }

    insert() {
        try {
            if (this.agentUID === undefined) { throw "No Agent UID assigned"; }
            const query = `INSERT INTO servers.ProcessorLogs (agentId, cpuIndex, vendorId, family, numberOfCores, modelName, speed, currentCPUUtilization)
            VALUES ('${this.agentUID}', ${this.cpuIndex}, '${this.vendorId}', '${this.family}', ${this.numberOfCores}, '${this.modelName}', '${this.speed}', '${JSON.stringify(this.currentCPUUtilization)}');`;
            return super.execQuery(query);
        } catch (err) {
            return err;
        }
    }

}
