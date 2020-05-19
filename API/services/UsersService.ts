import User from '../models/User';

module UsersService{
    export function getAll() {
        return null;
    }
    export async function create(agentUID: string, data: User) {
        const user = new User(data.activeUsers)
        user.setAgentUID(agentUID);
        return await user.insert();
    }
}

export default UsersService;
