import OperativeSystem from '../models/OperativeSystem';

module OperativeSystemsService{
    export function getAll() {
        return null;
    }
    export async function create(agentUID: string, data: OperativeSystem) {
        const operativeSystem = new OperativeSystem(data.runtime,data.name, data.platform);
        operativeSystem.setAgentUID(agentUID);
        return await operativeSystem.insert();
    }
}

export default OperativeSystemsService;
