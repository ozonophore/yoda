export const openSidebar = () => {
    if (typeof document !== 'undefined') {
        document.body.style.overflow = 'hidden';
        document.documentElement.style.setProperty('--SideNavigation-slideIn', '1');
    }
};

export const closeSidebar = () => {
    if (typeof document !== 'undefined') {
        document.documentElement.style.removeProperty('--SideNavigation-slideIn');
        document.body.style.removeProperty('overflow');
    }
};

export const toggleSidebar = () => {
    if (typeof window !== 'undefined' && typeof document !== 'undefined') {
        const slideIn = window
            .getComputedStyle(document.documentElement)
            .getPropertyValue('--SideNavigation-slideIn');
        if (slideIn) {
            closeSidebar();
        } else {
            openSidebar();
        }
    }
};

export function getValue<T, K extends keyof T>(data: T, key: K) {
    return data[key];
}

export const downloadFile = (url: string, fileName: string, onStartDownload: () => void, onFinishDownload: () => void, options?: RequestInit) => {
    onStartDownload()
    fetch(url, options)
        .then(res => res.blob())
        .then(blob => {
            const fileURL = URL.createObjectURL(blob);
            const fileLink = document.createElement('a');
            fileLink.href = fileURL;
            fileLink.download = fileName;
            fileLink.click();
        }).finally(() => onFinishDownload())
}