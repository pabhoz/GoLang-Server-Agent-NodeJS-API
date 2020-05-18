import { Request, Response } from "express"
import GodsService from "../services/GodsService"

/**
 * Retorna todos los dioses de la base de datos
 * 
 * @param req: Request => petición de express
 * @param res: Response => respuesta de express
 */
export async function gods(req: Request, res: Response) {
    const gods: any = await GodsService.getAll();
    res.status(200).json(gods);
}

export function create(req: Request, res: Response) {
    const gods = GodsService.create(req.body.name, req.body.origin);
    res.status(200).json(gods);
}

export function update(req: Request, res: Response) {
    
}

export function del(req: Request, res: Response) {
    
}
