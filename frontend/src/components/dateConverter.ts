export default function DateConverter(dateString: string): string {
    return new Date(dateString).toLocaleDateString('en-GB'); // DD/MM/YYYY
}