import { OpenAPI } from 'client';
import { downloadFile } from 'utils';

export class CustomSalesService {
    public static getSalesByMonthReport(
        onStartDownload: () => void,
        onFinishDownload: () => void,
        year: number,
        month: number,
    ) {

        const url = OpenAPI.BASE + `/sales/report?year=${year}&month=${month}`
        console.log("# " + url)
        const authHeader = `Bearer ${OpenAPI.TOKEN}`

        const options = {
            method: 'GET',
            headers: {
                Authorization: authHeader
            }
        };

        downloadFile(url,
            `sale_${year}-${month}.xlsx`,
            onStartDownload,
            onFinishDownload,
            options)
    }
}