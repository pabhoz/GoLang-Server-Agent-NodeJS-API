import Processor from '../models/Processor';

module ProcessorsService{
    export function getAll() {
        return null;
    }
    export async function create(agentUID: string, data: Processor) {
        console.log("Agent Name:", agentUID);
        const processor = new Processor(data.cpuIndex, data.vendorId, data.family, data.numberOfCores, data.modelName, data.speed, data.currentCPUUtilization);
        processor.setAgentUID(agentUID);
        console.log(processor);
        return await processor.insert();
    }
}

export default ProcessorsService;
