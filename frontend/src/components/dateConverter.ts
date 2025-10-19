import { parseISO, format } from 'date-fns';

export default function DateConverter(dateString: string): string {
    const d = parseISO(dateString);
    return format(d, 'dd/MM/yyyy');
}