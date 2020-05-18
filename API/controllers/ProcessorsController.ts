import { Request, Response } from "express";
import ProcessorsService from "../services/ProcessorsService";

export function getAll(req: Request, res: Response) {
    const gods = ProcessorsService.getAll();
    res.status(200).json(gods);
};

export async function createLog(req: Request, res: Response) {
    const agentUID = req.body.agentUID;
    const data = req.body.data;
    const result = await ProcessorsService.create(agentUID, data);
    
    const response: any = {
        error: false,
        msg: "Se ha creado un nuevo registro",
        result
    };
    if (typeof result != 'object') {
        response.error = true;
        response.msg = "Error al crear log de procesador";
    }
    res.status(200).json(response);
}
