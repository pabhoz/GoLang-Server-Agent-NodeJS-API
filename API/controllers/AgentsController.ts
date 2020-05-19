import { Request, Response } from "express";
import AgentsService from "../services/AgentsService";

export async function handshake(req: Request, res: Response) {

    const agentUID = req.body.agentUID;

    const exists = await AgentsService.get(agentUID);
    if (!exists) {
        const result = await AgentsService.create(agentUID);
        const response: any = {
            error: false,
            msg: "Se ha creado un nuevo registro",
            result
        };
        if (typeof result != 'object') {
            response.error = true;
            response.msg = "Error al crear log de procesos";
        }
        res.status(200).json(response);
    } else {
        res.status(200).json({
            error: false,
            msg: "Agente ya registrado",
        });
    }
}
