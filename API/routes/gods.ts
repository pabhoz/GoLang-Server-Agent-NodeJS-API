import express from "express";
import * as GodsController from "../controllers/GodsController";

module godsRouter {
    export function getRoutes() {
        let router: any = express.Router();
        router.route('/')
            .get('/', GodsController.gods)
            .post('/', GodsController.create)
            .put('/', GodsController.update)
            .delete('/', GodsController.del);
        return router;
    }
}
export default godsRouter;




