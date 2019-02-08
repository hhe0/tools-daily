import PyPDF2

path = 'C:\\Users\\LYPC\\Desktop\\简历.pdf'
pdfWriter = PyPDF2.PdfFileWriter()

pdfFileObj = open(path, 'rb')
pdfReader = PyPDF2.PdfFileReader(pdfFileObj)

for pageNum in range(0, pdfReader.numPages - 1):
    pageObj = pdfReader.getPage(pageNum)
    pdfWriter.addPage(pageObj)

pdfOutput = open('C:\\Users\\LYPC\\Desktop\\resume.pdf', 'wb')
pdfWriter.write(pdfOutput)
pdfOutput.close()







