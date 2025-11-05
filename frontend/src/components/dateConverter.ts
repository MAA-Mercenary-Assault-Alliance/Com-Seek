import { parseISO, format, isValid } from 'date-fns';

export default function DateConverter(dateString: string): string {
    const d = parseISO(dateString);
    if (!dateString || !isValid(d)) return 'Invalid Date';  // fallback for invalid dates
    return format(d, 'dd/MM/yyyy');
}