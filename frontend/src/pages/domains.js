import React, { useMemo, useState, useEffect, spacing } from 'react';
import axios from 'axios';
import {
  MRT_EditActionButtons,
  MaterialReactTable,
  useMaterialReactTable,
} from 'material-react-table';
import {
  Box,
  Button,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton,
  Tooltip,
} from '@mui/material';
import {
  QueryClient,
  QueryClientProvider,
  useMutation,
  useQuery,
  useQueryClient,
} from '@tanstack/react-query';
//import { fakeData, usStates } from './makeData';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import FileDownloadIcon from '@mui/icons-material/FileDownload';
import { mkConfig, generateCsv, download } from 'export-to-csv';

const csvConfig = mkConfig({
  fieldSeparator: ',',
  decimalSeparator: '.',
  useKeysAsHeaders: true,
});

const DomainView = () => {
  const [validationErrors, setValidationErrors] = useState({});

  const handleExportRows = (rows) => {
    const rowData = rows.map((row) => row.original);
    const csv = generateCsv(csvConfig)(rowData);
    download(csvConfig)(csv);
  };

  const handleExportData = (rows) => {
    //const csv = generateCsv(csvConfig)(fakeData);
    //download(csvConfig)(csv);
    const rowData = rows.map((row) => row.original);
    const csv = generateCsv(csvConfig)(rowData);
    download(csvConfig)(csv);
  };

  useEffect(() => {
    document.title = 'Domains';
  }, []);

  const columns = useMemo(
    () => [
      {
        accessorKey: 'ID',
        header: 'ID',
        enableEditing: false,
        size: 80,
      },
      {
        accessorKey: 'target',
        header: 'Target',
        muiEditTextFieldProps: {
          required: true,
          error: !!validationErrors?.target,
          helperText: validationErrors?.target,
          onFocus: () =>
            setValidationErrors({
              ...validationErrors,
              target: undefined,
            }),
        },
      },
      {
        accessorKey: 'domain',
        header: 'Domain',
        muiEditTextFieldProps: {
          required: true,
          error: !!validationErrors?.domain,
          helperText: validationErrors?.domain,
          onFocus: () =>
            setValidationErrors({
              ...validationErrors,
              domain: undefined,
            }),
        },
      },
      //{
      //  accessorKey: 'state',
      //  header: 'State',
      //  editVariant: 'select',
      //  editSelectOptions: usStates,
      //  muiEditTextFieldProps: {
      //    select: true,
      //    error: !!validationErrors?.state,
      //    helperText: validationErrors?.state,
      //  },
      //},
    ],
    [validationErrors],
  );

  //call CREATE hook
  const { mutateAsync: createDomain, isPending: isCreatingDomain } =
    useCreateDomain();
  //call READ hook
  const {
    data: fetchedDomains = [],
    isError: isLoadingDomainsError,
    isFetching: isFetchingDomains,
    isLoading: isLoadingDomains,
  } = useGetDomains();
  //call UPDATE hook
  const { mutateAsync: updateDomain, isPending: isUpdatingDomain } =
    useUpdateDomain();
  //call DELETE hook
  const { mutateAsync: deleteDomain, isPending: isDeletingDomain } =
    useDeleteDomain();

  //CREATE action
  const handleCreateDomain = async ({ values, table }) => {
    const newValidationErrors = validateDomain(values);
    if (Object.values(newValidationErrors).some((error) => error)) {
      setValidationErrors(newValidationErrors);
      return;
    }
    setValidationErrors({});
    await createDomain(values);
    table.setCreatingRow(null); //exit creating mode
  };

  //UPDATE action
  const handleSaveDomain = async ({ values, table }) => {
    const newValidationErrors = validateDomain(values);
    if (Object.values(newValidationErrors).some((error) => error)) {
      setValidationErrors(newValidationErrors);
      return;
    }
    setValidationErrors({});
    await updateDomain(values);
    table.setEditingRow(null); //exit editing mode
  };

  //DELETE action
  const openDeleteConfirmModal = (row) => {
    if (window.confirm('Are you sure you want to delete this record [' + row.original.ID + ']?')) {
      deleteDomain(row.original.ID);
    }
  };

  const table = useMaterialReactTable({
    //initialState: { columnVisibility: { ID: false } },
    initialState: {
      columnVisibility: { ID: false }
    },
    columns,
    data: fetchedDomains,
    enableRowSelection: true,
    columnFilterDisplayMode: 'popover',
    createDisplayMode: 'modal', //default ('row', and 'custom' are also available)
    editDisplayMode: 'modal', //default ('row', 'cell', 'table', and 'custom' are also available)
    enableEditing: true,
    getRowId: (row) => row.id,
    muiToolbarAlertBannerProps: isLoadingDomainsError
      ? {
        color: 'error',
        children: 'Error loading data',
      }
      : undefined,
    muiTableContainerProps: {
      sx: {
        minHeight: '500px',
      },
    },
    onCreatingRowCancel: () => setValidationErrors({}),
    onCreatingRowSave: handleCreateDomain,
    onEditingRowCancel: () => setValidationErrors({}),
    onEditingRowSave: handleSaveDomain,
    //optionally customize modal content
    renderCreateRowDialogContent: ({ table, row, internalEditComponents }) => (
      <>
        <DialogTitle variant="p">Create New Domain</DialogTitle>
        <DialogContent
          sx={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}
        >
          {internalEditComponents} {/* or render custom edit components here */}
        </DialogContent>
        <DialogActions>
          <MRT_EditActionButtons variant="text" table={table} row={row} />
        </DialogActions>
      </>
    ),
    //optionally customize modal content
    renderEditRowDialogContent: ({ table, row, internalEditComponents }) => (
      <>
        <DialogTitle variant="p">Edit Domain</DialogTitle>
        <DialogContent
          sx={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}
        >
          {internalEditComponents} {/* or render custom edit components here */}
        </DialogContent>
        <DialogActions>
          <MRT_EditActionButtons variant="text" table={table} row={row} />
        </DialogActions>
      </>
    ),
    renderRowActions: ({ row, table }) => (
      <Box sx={{ display: 'flex', gap: '1rem', width: '100px' }}>
        <Tooltip title="Edit">
          <IconButton onClick={() => table.setEditingRow(row)}>
            <EditIcon />
          </IconButton>
        </Tooltip>
        <Tooltip title="Delete">
          <IconButton color="error" onClick={() => openDeleteConfirmModal(row)}>
            <DeleteIcon />
          </IconButton>
        </Tooltip>
      </Box>
    ),
    renderTopToolbarCustomActions: ({ table }) => (
      <Box
        sx={{
          display: 'flex',
          gap: '16px',
          padding: '8px',
          flexWrap: 'wrap',
        }}
      >
        <Button
          variant="contained"
          onClick={() => {
            table.setCreatingRow(true); //simplest way to open the create row modal with no default values
            //or you can pass in a row object to set default values with the `createRow` helper function
            // table.setCreatingRow(
            //   createRow(table, {
            //     //optionally pass in default values for the new row, useful for nested data or other complex scenarios
            //   }),
            // );
          }}
        >
          Create New Domain
        </Button>
        <Button
          disabled={table.getPrePaginationRowModel().rows.length === 0}
          //export all rows, including from the next page, (still respects filtering and sorting)
          onClick={() =>
            handleExportRows(table.getPrePaginationRowModel().rows)
          }
          startIcon={<FileDownloadIcon />}
        >
          Export All Rows
        </Button>
        <Button
          disabled={table.getRowModel().rows.length === 0}
          //export all rows as seen on the screen (respects pagination, sorting, filtering, etc.)
          onClick={() => handleExportRows(table.getRowModel().rows)}
          startIcon={<FileDownloadIcon />}
        >
          Export Page Rows
        </Button>
        <Button
          disabled={
            !table.getIsSomeRowsSelected() && !table.getIsAllRowsSelected()
          }
          //only export selected rows
          onClick={() => handleExportRows(table.getSelectedRowModel().rows)}
          startIcon={<FileDownloadIcon />}
        >
          Export Selected Rows
        </Button>
      </Box>
    ),
    state: {
      isLoading: isLoadingDomains,
      isSaving: isCreatingDomain || isUpdatingDomain || isDeletingDomain,
      showAlertBanner: isLoadingDomainsError,
      showProgressBars: isFetchingDomains,
    },
  });

  return <MaterialReactTable table={table} />;
};

//CREATE hook (post new record to api)
function useCreateDomain() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (record) => {
      console.log(record);
      try {
        record['ID'] = 0;
        const data = record;
        const response = await axios.post('http://localhost:8000/api/domain/add', data);
        console.log('Domain saved:', response.data);
        return response;
      } catch (error) {
        console.error('Error saving domain:', error);
      }
    },
    //client side optimistic update
    //onMutate: (newDomainInfo) => {
    //  queryClient.setQueryData(['domains'], (prevDomains) => [
    //    ...prevDomains,
    //    {
    //      ...newDomainInfo,
    //      id: (Math.random() + 1).toString(36).substring(7),
    //    },
    //  ]);
    //},
    onSettled: () => queryClient.invalidateQueries({ queryKey: ['domains'] }), //refetch domains after mutation, disabled for demo
  });
}

//READ hook (get domains from api)
function useGetDomains() {

  //var domains = [];

  //var domain = {
  //  id: 0,
  //  target: "faketarget",
  //  domain: "staticdomain.com"
  //};

  //domains.push(domain);

  return useQuery({
    queryKey: ['domains'],
    queryFn: async () => {
      const response = await fetch('http://localhost:8000/api/domains/');
      const json = await response.json();
      return json;

    },
    refetchOnWindowFocus: false,
  });
}

//UPDATE hook (put user in api)
function useUpdateDomain() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (row) => {
      console.log("Domain ID:" + row.ID);
      try {
        const response = await axios.put('http://localhost:8000/api/domain/update/' + row.ID, row);
        console.log('Domain updated:', response.data);
        return response;
      } catch (error) {
        console.error('Error updating domain:', error);
      }

    },
    //client side optimistic update
    onMutate: (newDomainInfo) => {
      queryClient.setQueryData(['domains'], (prevDomains) =>
        prevDomains?.map((prevDomain) =>
          prevDomain.id === newDomainInfo.id ? newDomainInfo : prevDomain,
        ),
      );
    },
    onSettled: () => queryClient.invalidateQueries({ queryKey: ['domains'] }), //refetch domains after mutation, disabled for demo
  });
}

//DELETE hook (delete user in api)
function useDeleteDomain() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (domainId) => {
      //send api update request here
      //await new Promise((resolve) => setTimeout(resolve, 1000)); //fake api call
      //return Promise.resolve();
      console.log("Domain ID:" + domainId);
      try {
        const response = await axios.delete('http://localhost:8000/api/domain/delete/' + domainId);
        console.log('Domain deleted:', response.data);
        return response;
      } catch (error) {
        console.error('Error deleting domain:', error);
      }
    },
    // client side optimistic update
    onMutate: (domainId) => {
      queryClient.setQueryData(['domains'], (prevDomains) =>
        prevDomains?.filter((domain) => domain.id !== domainId),
      );
    },
    onSettled: () => queryClient.invalidateQueries({ queryKey: ['domains'] }), //refetch domains after mutation, disabled for demo
  });
}

const queryClient = new QueryClient();

const ExampleWithProviders = () => (
  //Put this with your other react-query providers near root of your app
  <QueryClientProvider client={queryClient}>
    <DomainView />
  </QueryClientProvider>
);

export default ExampleWithProviders;

const validateRequired = (value) => !!value.length;
//const validateEmail = (email) =>
//  !!email.length &&
//  email
//    .toLowerCase()
//    .match(
//      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
//    );

function validateDomain(record) {
  return {
    target: !validateRequired(record.target)
      ? 'Target is Required'
      : '',
    domain: !validateRequired(record.domain) 
      ? 'Domain is Required' 
      : '',
    //email: !validateEmail(user.email) ? 'Incorrect Email Format' : '',
  };
}
