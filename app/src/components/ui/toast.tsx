import toast from 'react-hot-toast';

export function showSuccessToastMessage(message: string): void {
    toast.success(<b>{message}</b>);
}

export function showErrorToastMessage(message: string): void {
    toast.error(<b>{message}</b>);
}

export function showInfoToastMessage(message: string): void {
    toast(<b>{message}</b>, {
        position: 'bottom-right',
    });
}
