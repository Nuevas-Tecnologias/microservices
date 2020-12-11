import express, {Router} from 'express';
import helmet from 'helmet';
import bodyParser from 'body-parser';
import { Consumer } from 'sqs-consumer';
import { saveTechOrder } from "./saveTechOrder";
import { getOrder } from "./techOrdersApi";

const port = parseInt(process.env.PORT || 3000, 10);

const AWS = require('aws-sdk');

AWS.config.update({region: 'us-west-2'});

const app = express();

app.use(bodyParser.json());
app.use(helmet());

const routes = Router();

app.use('/orders', routes);

routes.get('/:id', async (req, res, next) => {
    const httpResponse = await getOrder(req.params.id);
    res.status(200).json(httpResponse);
});

const sqsListener = Consumer.create({
    queueUrl: 'https://sqs.us-west-2.amazonaws.com/881619806726/save-tech-order-command.fifo',
    handleMessage: async (message) => {
        console.log('Processing message');
        await saveTechOrder(message.Body);
        console.log('Message processed');
    }
});

sqsListener.start();

app.listen(port, () => {
    // eslint-disable-next-line no-console
    console.log(`Running HTTP Server on port: ${port}`);
});
