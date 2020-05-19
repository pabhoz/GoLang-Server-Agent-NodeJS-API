import { Request, Response } from "express";
import UsersService from "../services/UsersService";

export async function createLog(req: Request, res: Response) {
    const agentUID = req.body.agentUID;
    const data = req.body.data;
    const result = await UsersService.create(agentUID, data);
    
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
}
