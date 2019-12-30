import PyPDF2
import os

path = 'C:\\Users\LYPC\Desktop\新建文件夹\\'
sum = 0
pdfWriter = PyPDF2.PdfFileWriter()
for file in os.listdir(path):
    pdfFileObj = open(path + file, 'rb')
    pdfReader = PyPDF2.PdfFileReader(pdfFileObj)
    sum += pdfReader.numPages

    for pageNum in range(0, pdfReader.numPages):
        pageObj = pdfReader.getPage(pageNum)
        pdfWriter.addPage(pageObj)


# 计算该文件夹下所有PDF文件的页数和
print(sum)

# 将多个PDF合并为一个PDF文件
pdfOutput = open('C:\\Users\LYPC\Desktop\merge.pdf', 'wb')
pdfWriter.write(pdfOutput)
pdfOutput.close()



