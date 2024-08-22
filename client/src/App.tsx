import React, { useState, useEffect } from 'react';
import { Moon, Sun, Database, Plus, Trash2 } from 'lucide-react';
import { Table, TableHeader, TableBody, TableRow, TableHead, TableCell } from '@/components/ui/table';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { ScrollArea } from '@/components/ui/scroll-area';
import AxiosClient from './lib/http';

type Model = string;
type Record = { [key: string]: any };

const App: React.FC = () => {
  const [models, setModels] = useState<Model[]>(["Users"]);
  const [selectedModel, setSelectedModel] = useState<Model>('');
  const [records, setRecords] = useState<Record[]>([]);
  const [newRecord, setNewRecord] = useState<Record>({});
  const [theme, setTheme] = useState<'light' | 'dark'>('light');

  useEffect(() => {
    fetchModels();
  }, []);

  useEffect(() => {
    if (selectedModel) {
      fetchRecords(selectedModel);
    }
  }, [selectedModel]);

  useEffect(() => {
    document.body.className = theme;
  }, [theme]);

  const fetchModels = async (): Promise<void> => {
    const res = await AxiosClient.get("/health");
    console.log(res);
  };

  const fetchRecords = async (model: Model): Promise<void> => {
    setRecords([
      { id: 1, name: 'John Doe', email: 'john@example.com' },
      { id: 2, name: 'Jane Doe', email: 'jane@example.com' },
    ]);
  };

  const handleModelChange = (model: Model): void => {
    setSelectedModel(model);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setNewRecord({ ...newRecord, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    fetchRecords(selectedModel);
    setNewRecord({});
  };

  const toggleTheme = (): void => {
    setTheme(theme === 'dark' ? 'light' : 'dark');
  };

  return (
    <div className={`flex h-screen bg-${theme === 'dark' ? 'gray-800' : 'white'} text-${theme === 'dark' ? 'white' : 'black'}`}>
      <div className={`w-64 border-r border-${theme === 'dark' ? 'gray-700' : 'gray-200'} p-4`}>
        <div className="flex items-center justify-between mb-6">
          <h1 className="text-2xl font-bold">GORM Studio</h1>
          <Button variant="ghost" size="icon" onClick={toggleTheme}>
            {theme === 'dark' ? <Sun className="h-[1.2rem] w-[1.2rem]" /> : <Moon className="h-[1.2rem] w-[1.2rem]" />}
          </Button>
        </div>
        <nav>
          {models.map((model) => (
            <Button
              key={model}
              variant={selectedModel === model ? 'secondary' : 'ghost'}
              className="w-full justify-start mb-2"
              onClick={() => handleModelChange(model)}
            >
              <Database className="mr-2 h-4 w-4" />
              {model}
            </Button>
          ))}
        </nav>
      </div>
      
      <div className="flex-1 flex flex-col">
        <div className={`border-b border-${theme === 'dark' ? 'gray-700' : 'gray-200'} p-4 flex justify-between items-center`}>
          <h2 className="text-xl font-semibold">{selectedModel || 'Select a model'}</h2>
          <Dialog>
            <DialogTrigger asChild>
              <Button>
                <Plus className="mr-2 h-4 w-4" />
                Add record
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Add new {selectedModel} record</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleSubmit} className="space-y-4">
                {records.length > 0 && Object.keys(records[0]).map((key) => (
                  <Input
                    key={key}
                    name={key}
                    placeholder={key}
                    value={newRecord[key] || ''}
                    onChange={handleInputChange}
                  />
                ))}
                <Button type="submit">Add Record</Button>
              </form>
            </DialogContent>
          </Dialog>
        </div>

        <ScrollArea className="flex-1 p-4">
          <Table>
            <TableHeader>
              <TableRow>
                {records.length > 0 && Object.keys(records[0]).map((key) => (
                  <TableHead key={key}>{key}</TableHead>
                ))}
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {records.map((record, index) => (
                <TableRow key={index}>
                  {Object.values(record).map((value, i) => (
                    <TableCell key={i}>{value}</TableCell>
                  ))}
                  <TableCell>
                    <Button variant="ghost" size="icon">
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </ScrollArea>
      </div>
    </div>
  );
};

export default App;
