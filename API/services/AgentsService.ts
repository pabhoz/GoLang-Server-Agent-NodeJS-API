import Agent from '../models/Agent';

module AgentsService{
    export function getAll() {
        return null;
    }
    export async function create(agentUID: string) {
        const agent = new Agent(agentUID)
        return await agent.insert();
    }
    export async function get(agentUID: string) {
        const agent = new Agent(agentUID)
        const theOne = await agent.findOne();
        return (theOne.length > 0) ? theOne[0] : undefined;
    }
}

export default AgentsService;
