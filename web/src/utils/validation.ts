import * as yup from 'yup'

export const aiAnalysisSchema = yup.object({
  ticker: yup
    .string()
    .required('Ticker is required')
    .min(1, 'Ticker cannot be empty')
    .max(10, 'Ticker is too long')
    .matches(/^[A-Za-z.]+$/, 'Ticker must contain only letters and dots'),
  company: yup
    .string()
    .required('Company name is required')
    .min(2, 'Company name is too short')
    .max(100, 'Company name is too long')
})

export type AIAnalysisFormData = yup.InferType<typeof aiAnalysisSchema>
