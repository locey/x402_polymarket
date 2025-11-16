import Footer from '@/components/Footer';
import Header from '@/components/Header';

const UIProvider = ({ children }: { children: React.ReactNode }) => {
  return (
    <div>
      <Header />
      <main>{children}</main>
      <Footer />
    </div>
  );
};

export default UIProvider;
