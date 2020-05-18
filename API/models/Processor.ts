import Model from '../libs/Model';

interface CPUUtilizationI {
    index: number,
    utilization: string
}

export default class Processor extends Model {

    static dbc: any;

    index: number;
    vendorId: string;
    family: string;
    numberOfCores: number;
    modelName: string;
    speed: string;
    currentCPUUtilization: Array<CPUUtilizationI>

    constructor(index: number,
        vendorId: string,
        family: string,
        numberOfCores: number,
        modelName: string,
        speed: string,
        currentCPUUtilization: Array<CPUUtilizationI>) {
        super();
        this.index = index;
        this.vendorId = vendorId;
        this.family = family;
        this.numberOfCores = numberOfCores;
        this.modelName = modelName;
        this.speed = speed;
        this.currentCPUUtilization = currentCPUUtilization;
    }

}
