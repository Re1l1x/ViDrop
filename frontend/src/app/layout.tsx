import type { Metadata } from "next";
import { Montserrat, Audiowide } from "next/font/google";
import "./globals.css";

const montserrat = Montserrat({ subsets: ["latin"] });
const audiowide = Audiowide({ weight: "400", subsets: ["latin"] });

export const metadata: Metadata = {
    title: "ViDrop",
    description: "ViDrop",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en" className={montserrat.className}>
            <body>{children}</body>
        </html>
    );
}
