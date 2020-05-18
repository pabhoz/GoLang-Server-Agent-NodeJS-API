import Processor from '../models/Processor';

module ProcessorsService{
    export function getAll() {
        return null;
    }
    export function create(agentName: string, data: Processor) {
        console.log("Agent Name:", agentName);
        const processor = new Processor(data.index, data.vendorId, data.family, data.numberOfCores, data.modelName, data.speed, data.currentCPUUtilization);
        console.log("Processor:", processor);
        return processor;
    }
}

export default ProcessorsService;
