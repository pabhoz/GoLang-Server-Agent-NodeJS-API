import RunningProcesses from '../models/RunningProcesses';

module RunningProcessesService{
    export function getAll() {
        return null;
    }
    export async function create(agentUID: string, data: RunningProcesses) {
        const runningProcesses = new RunningProcesses(data.total, data.running, data.processesList)
        runningProcesses.setAgentUID(agentUID);
        return await runningProcesses.insert();
    }
}

export default RunningProcessesService;
